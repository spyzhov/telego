#!/usr/bin/python3.7
import logging
import requests
import re

from bs4 import BeautifulSoup
from bs4.element import Tag

from dataclasses import dataclass
from typing import List, Optional, Generator, Union, Tuple

URL = r'https://core.telegram.org/bots/api'

TYPES = {
    'String': 'string',
    'Integer': 'int',
    'Int': 'int',
    'Float number': 'float64',
    'Float': 'float64',
    'Boolean': 'bool',
    'True': 'bool',
}


@dataclass
class Field:
    name: str
    type: str
    description: str
    required: bool

    def __str__(self):
        name = ''.join([name.capitalize() for name in self.name.split('_')])
        _type = get_type(self.type)
        omitempty = ',omitempty' if not self.required else ''
        return f'\t// {self.description}\n\t{name:20} {_type:20} `json:"{self.name}{omitempty}"`'

    def __repr__(self):
        return f'//    {self.name:20} {self.type:20} {self.description}'


@dataclass
class Type:
    name: str
    url: str
    description: str
    fields: List[Field]
    postfix: str

    @property
    def uname(self) -> str:
        return self.name[0].upper() + self.name[1:]

    @property
    def exists(self) -> bool:
        return len(self.fields) > 0

    def __str__(self):
        parts = ['type %s%s struct {' % (self.uname, self.postfix)]
        for field in self.fields:
            parts.append(f'{field}')
        parts.append('}')
        return '\n'.join(parts)

    def __repr__(self):
        parts = [
            f'// {self.name}',
            f'// {self.url}',
            '//',
            f'{description(self.description)}',
            '//',
        ]
        if len(self.fields) > 0:
            parts.append('//    %-20s %-20s %s' % ('Field', 'Type', 'Description'))
            for field in self.fields:
                parts.append(f'{field!r}')
            parts.append('//')
        return '\n'.join(parts)


def find(soup: Union[BeautifulSoup, Tag], *args) -> Optional[Tag]:
    return soup.find(*args)


def find_all(soup: Union[BeautifulSoup, Tag], *args) -> List[Tag]:
    return soup.find_all(*args)


def parent(tag: Tag) -> Optional[Tag]:
    return tag.parent


def next_gen(tag: Tag) -> Generator[Tag, None, None]:
    return tag.next_elements


def next(tag: Tag) -> Optional[Tag]:
    return tag.next


def get_text(tag: Tag) -> str:
    return tag.get_text().strip()


def get_type(base: str) -> str:
    if base.startswith('Array of '):
        return '[]' + get_type(base[9:])
    if base.find(' or ') != -1:
        base = base.split(' or ')[0]
    if base.find(' and ') != -1:
        base = base.split(' and ')[0]
    return TYPES.get(base, '*' + base)


def parse_responses(soup: BeautifulSoup) -> Tuple[List[Type], List[Type]]:
    request = list()
    response = list()
    h3 = parent(find(soup, 'a', {'name': 'getting-updates'}))
    current = None  # Type
    for tag in next_gen(h3):
        if tag.name == 'h3':
            logging.debug("found Level: " + tag.text)
            continue
        if tag.name == 'h4':
            current = None  # Type
            logging.debug("found H4: " + tag.text)
            name = get_text(tag)

            if name.find(' ') != -1:
                logging.debug("skip title: " + tag.text)
                continue

            anchor = find(tag, 'a', {'class': 'anchor'}).get_attribute_list("href", [''])[0]
            current = Type(get_text(tag), URL + anchor, "", list(), "")

            if name[0].isupper():
                logging.debug("found response: " + tag.text)
                response.append(current)
            else:
                logging.debug("found request: " + tag.text)
                current.postfix = 'Request'
                request.append(current)

        elif current is not None and tag.name == 'p' and current.description == "":
            current.description = get_text(tag).replace("\n", " ")
            logging.debug("found Description: " + current.description)
        elif current is not None and tag.name == 'table':
            tbody = find(tag, 'tbody')
            for tr in find_all(tbody, 'tr'):
                td = find_all(tr, 'td')
                field = Field(get_text(td[0]), get_text(td[1]), get_text(td[-1]).replace("\n", " "), True)
                if len(td) == 4 and get_text(td[2]) == 'Optional':
                    field.description = 'Optional. ' + field.description
                field.required = not field.description.startswith('Optional.')
                current.fields.append(field)
    return request, response


def get_result(desc: str) -> str:
    m = None
    if m is None:
        m = re.search('Returns basic information about the bot in form of a ([\w\s]+) object', desc)
    if m is None:
        m = re.search('On success, if the edited message was sent by the bot, the edited (\w+) is returned, otherwise True is returned', desc)
    if m is None:
        m = re.search('On success, if edited message is sent by the bot, the edited (\w+) is returned, otherwise True is returned', desc)
    if m is None:
        m = re.search('On success, if the message was sent by the bot, the sent (\w+) is returned, otherwise True is returned', desc)
    if m is None:
        m = re.search('On success, if the message was sent by the bot, returns the edited (\w+), otherwise returns True', desc)
    if m is None:
        m = re.search('On success, the stopped (\w+) with the final results is returned', desc)
    if m is None:
        m = re.search('On success, an array of the sent (\w+)s is returned', desc)
        if m is not None:
            return get_type('Array of '+m.group(1))
    if m is None:
        m = re.search('Returns the new invite link as ([\w\s]+) on success', desc)
    if m is None:
        m = re.search('Returns the uploaded (\w+) on success', desc)
    if m is None:
        m = re.search('Returns an? ([\w\s]+) objects? on success', desc)
    if m is None:
        m = re.search('Returns ([\w\s]+) on success', desc)
    if m is None:
        m = re.search('On success, the sent ([\w\s]+) is returned', desc)
    if m is None:
        m = re.search('On success, returns an? ([\w\s]+) objects?', desc)
    if m is None:
        m = re.search('On success, a ([\w\s]+) object is returned', desc)
    if m is None:
        m = re.search('On success, ([\w\s]+) is returned', desc)
    if m is None:
        m = re.search('An ([\w\s]+) objects is returned', desc)
    if m is not None:
        return get_type(m.group(1))
    return 'interface{}'


def description(text: str) -> str:
    total = 0
    result = []
    for word in filter(lambda x: len(x) > 0, text.split(' ')):
        total += len(word)+1
        index = total//100
        try:
            result[index]
        except IndexError:
            result.append('//')
        finally:
            result[index] += ' ' + word
    return '\n'.join(result)


def method(_type: Type) -> str:
    result_type = get_result(_type.description)
    result = ''
    if result_type.startswith('[]'):
        result = f'\n\tresult = make({result_type}, 0)'
    elif result_type.startswith('*'):
        result = f'\n\tresult = new({result_type[1:]})'
    request = f', request *{_type.uname}{_type.postfix}' if _type.exists else ''
    nil = 'request' if _type.exists else 'nil'
    return f'''// {_type.name}
// {_type.url}
{description(_type.description)}
func (b *Bot) {_type.uname}(ctx context.Context{request}) (result {result_type}, err error) {{{result}
\treturn result, b.postResult(ctx, "{_type.name}", {nil}, &result)
}}
'''


def main():
    logging.basicConfig(
        level=logging.DEBUG,
        format='%(asctime)-15s %(levelname)s %(name)-8s %(message)s',
    )
    content = requests.get(URL).text
    soup = BeautifulSoup(content, 'html.parser')
    request, response = parse_responses(soup)

    with open('response.go', 'w') as objects:
        objects.write('package telego\n\n')
        for _type in response:
            objects.write(f'{_type!r}\n{_type}\n\n')

    with open('request.go', 'w') as objects:
        objects.write('package telego\n\n')
        for _type in request:
            if _type.exists:
                objects.write(f'{_type!r}\n{_type}\n\n')

    with open('methods.go', 'w') as objects:
        objects.write('package telego\n\nimport "context"\n\n')
        for _type in request:
            objects.write(f'{method(_type)}\n\n')


if __name__ == "__main__":
    main()
