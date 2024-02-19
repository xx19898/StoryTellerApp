export function tagToType(tag:string){
    if(tag === 'h2') return 'title'
    if(tag === 'p') return 'paragraph'
    if(tag === 'img') return 'image'
    return 'unknown'
}

export function typeToTag(type: 'title' | 'paragraph' | 'image'){
    if(type === 'title') return 'h2'
    if(type === 'paragraph') return 'p'
    if(type === 'image') return 'img'
    return 'unknown'
}