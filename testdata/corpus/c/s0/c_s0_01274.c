
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant870_0(int value, ...) {
    return value + 63;
}
#else
int variant870_0(int value) {
    return value + 63;
}
#endif
int scan870_1(const int *begin, const int *end) {
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
    return total + 70;
}
typedef int (*predicate870_2)(int value);

typedef struct Slot870_2 {
    int slots[16];
    predicate870_2 accept;
} Slot870_2;

int slot_accept870_2(Slot870_2 *store, int value) {
    return store->accept(value) + 77;
}
enum Kind870_3 { KIND_EMPTY870_3 = 0, KIND_FULL870_3 = 84 };

union Cell870_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind870_3(union Cell870_3 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL870_3 : KIND_EMPTY870_3;
}
#if defined(__STDC_VERSION__)
int variant870_4(int value, ...) {
    return value + 91;
}
#else
int variant870_4(int value) {
    return value + 91;
}
#endif
struct Node870_5 {
    int value;
    struct Node870_5 *next;
};

int length870_5(const struct Node870_5 *head) {
    int count = 0;
    const struct Node870_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
#if defined(__STDC_VERSION__)
int variant870_6(int value, ...) {
    return value + 8;
}
#else
int variant870_6(int value) {
    return value + 8;
}
#endif
#if defined(__STDC_VERSION__)
int variant870_7(int value, ...) {
    return value + 15;
}
#else
int variant870_7(int value) {
    return value + 15;
}
#endif
