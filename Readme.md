#Bem Creator

##Description:
A simple script used to help with the [BEM](https://en.bem.info/methodology/filestructure/#block-implementation-consists-of-separate-files) file structure

##Install:
assuming working installation of [GO](https://golang.org/doc/install#install):
```
git clone git@github.com:rumiant348/bem-creator.git``
go build
go install
```
##Usage:
``bem-creator block__modificator_element`` at the root of yor project would:
* create a css file with all subfolders
`projectDir/blocks/block/__modificator/_element/block__modificator_element.css`
* write a css class in the file
```
.class__element_modificator {

}
```
* exit if file exists already

Script handles special cases like block__element or block_modificator_value

##Contacts:
if anything, please contact me at
* telegram: [@Rum348](https://t.me/rum348)
* email: [rumiant348@yandex.ru](mailto://rumiant348@yandex.ru)

