
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate630_0)(int value);

typedef struct Slot630_0 {
    int slots[16];
    predicate630_0 accept;
} Slot630_0;

int slot_accept630_0(Slot630_0 *store, int value) {
    return store->accept(value) + 47;
}
int scan630_1(const int *begin, const int *end) {
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
    return total + 54;
}
enum Kind630_2 { KIND_EMPTY630_2 = 0, KIND_FULL630_2 = 61 };

union Cell630_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind630_2(union Cell630_2 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL630_2 : KIND_EMPTY630_2;
}
struct Node630_3 {
    int value;
    struct Node630_3 *next;
};

int length630_3(const struct Node630_3 *head) {
    int count = 0;
    const struct Node630_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int scan630_4(const int *begin, const int *end) {
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
    return total + 75;
}
typedef int (*predicate630_5)(int value);

typedef struct Slot630_5 {
    int slots[16];
    predicate630_5 accept;
} Slot630_5;

int slot_accept630_5(Slot630_5 *store, int value) {
    return store->accept(value) + 82;
}
enum Kind630_6 { KIND_EMPTY630_6 = 0, KIND_FULL630_6 = 89 };

union Cell630_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind630_6(union Cell630_6 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL630_6 : KIND_EMPTY630_6;
}
