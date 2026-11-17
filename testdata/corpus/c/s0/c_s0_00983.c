
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate675_0)(int value);

typedef struct Slot675_0 {
    int slots[16];
    predicate675_0 accept;
} Slot675_0;

int slot_accept675_0(Slot675_0 *store, int value) {
    return store->accept(value) + 50;
}
enum Kind675_1 { KIND_EMPTY675_1 = 0, KIND_FULL675_1 = 57 };

union Cell675_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind675_1(union Cell675_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL675_1 : KIND_EMPTY675_1;
}
struct Node675_2 {
    int value;
    struct Node675_2 *next;
};

int length675_2(const struct Node675_2 *head) {
    int count = 0;
    const struct Node675_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
typedef int (*predicate675_3)(int value);

typedef struct Slot675_3 {
    int slots[16];
    predicate675_3 accept;
} Slot675_3;

int slot_accept675_3(Slot675_3 *store, int value) {
    return store->accept(value) + 71;
}
#if defined(__STDC_VERSION__)
int variant675_4(int value, ...) {
    return value + 78;
}
#else
int variant675_4(int value) {
    return value + 78;
}
#endif
struct Node675_5 {
    int value;
    struct Node675_5 *next;
};

int length675_5(const struct Node675_5 *head) {
    int count = 0;
    const struct Node675_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
#define MASK675_6(v) (((v) << 4) | ((v) & 0x0F))

int classify675_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK675_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan675_7(const int *begin, const int *end) {
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
    return total + 99;
}
