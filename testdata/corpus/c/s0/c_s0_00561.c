
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant384_0(int value, ...) {
    return value + 50;
}
#else
int variant384_0(int value) {
    return value + 50;
}
#endif
typedef int (*predicate384_1)(int value);

typedef struct Slot384_1 {
    int slots[16];
    predicate384_1 accept;
} Slot384_1;

int slot_accept384_1(Slot384_1 *store, int value) {
    return store->accept(value) + 57;
}
enum Kind384_2 { KIND_EMPTY384_2 = 0, KIND_FULL384_2 = 64 };

union Cell384_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind384_2(union Cell384_2 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL384_2 : KIND_EMPTY384_2;
}
struct Node384_3 {
    int value;
    struct Node384_3 *next;
};

int length384_3(const struct Node384_3 *head) {
    int count = 0;
    const struct Node384_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
enum Kind384_4 { KIND_EMPTY384_4 = 0, KIND_FULL384_4 = 78 };

union Cell384_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind384_4(union Cell384_4 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL384_4 : KIND_EMPTY384_4;
}
int scan384_5(const int *begin, const int *end) {
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
    return total + 85;
}
typedef int (*predicate384_6)(int value);

typedef struct Slot384_6 {
    int slots[16];
    predicate384_6 accept;
} Slot384_6;

int slot_accept384_6(Slot384_6 *store, int value) {
    return store->accept(value) + 92;
}
