/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// cRecipe covers structs, unions, enums, typedefs, function pointers,
// preprocessor macros, bit operations, pointer arithmetic and loop forms.
func cRecipe() Recipe {
	return Recipe{
		Language: "c",
		Ext:      "c",
		Tags:     []string{"struct", "macro", "pointer", "bitops", "controlflow", "function-pointer", "enum"},
		Header:   "\n",
		AnchorClean: `int anchor_probe(void) {
    return 1 + 2 == 3;
}

`,
		AnchorFull: `int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

`,
		CleanBlock: "int step@@I@@_@@S@@(void) { return @@N@@ + @@S@@; }\n",
		Blocks: []BlockTemplate{
			{Text: `typedef int (*predicate@@I@@_@@S@@)(int value);

typedef struct Slot@@I@@_@@S@@ {
    int slots[16];
    predicate@@I@@_@@S@@ accept;
} Slot@@I@@_@@S@@;

int slot_accept@@I@@_@@S@@(Slot@@I@@_@@S@@ *store, int value) {
    return store->accept(value) + @@N@@;
}
`},
			{Text: `#define MASK@@I@@_@@S@@(v) (((v) << 4) | ((v) & 0x0F))

int classify@@I@@_@@S@@(int value) {
    switch (value & 0x0F) {
    case 0: return MASK@@I@@_@@S@@(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
`},
			{Text: `int scan@@I@@_@@S@@(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + @@N@@;
}
`},
			{Text: `enum Kind@@I@@_@@S@@ { KIND_EMPTY@@I@@_@@S@@ = 0, KIND_FULL@@I@@_@@S@@ = @@N@@ };

union Cell@@I@@_@@S@@ {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind@@I@@_@@S@@(union Cell@@I@@_@@S@@ *cell) {
    cell->number = @@M@@;
    return cell->number > 0 ? KIND_FULL@@I@@_@@S@@ : KIND_EMPTY@@I@@_@@S@@;
}
`},
			{Text: `#if defined(__STDC_VERSION__)
int variant@@I@@_@@S@@(int value, ...) {
    return value + @@N@@;
}
#else
int variant@@I@@_@@S@@(int value) {
    return value + @@N@@;
}
#endif
`},
			{Text: `int gcd@@I@@_@@S@@(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale@@I@@_@@S@@(int value) {
    return gcd@@I@@_@@S@@(value, @@M@@) * @@N@@;
}
`},
			{Text: `struct Node@@I@@_@@S@@ {
    int value;
    struct Node@@I@@_@@S@@ *next;
};

int length@@I@@_@@S@@(const struct Node@@I@@_@@S@@ *head) {
    int count = 0;
    const struct Node@@I@@_@@S@@ *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + @@N@@;
}
`},
		},
	}
}
