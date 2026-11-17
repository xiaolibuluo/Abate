
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node668_0 {
    int value;
    struct Node668_0 *next;
};

int length668_0(const struct Node668_0 *head) {
    int count = 0;
    const struct Node668_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
typedef int (*predicate668_1)(int value);

typedef struct Slot668_1 {
    int slots[16];
    predicate668_1 accept;
} Slot668_1;

int slot_accept668_1(Slot668_1 *store, int value) {
    return store->accept(value) + 63;
}
int scan668_2(const int *begin, const int *end) {
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
int scan668_3(const int *begin, const int *end) {
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
    return total + 77;
}
enum Kind668_4 { KIND_EMPTY668_4 = 0, KIND_FULL668_4 = 84 };

union Cell668_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind668_4(union Cell668_4 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL668_4 : KIND_EMPTY668_4;
}
struct Node668_5 {
    int value;
    struct Node668_5 *next;
};

int length668_5(const struct Node668_5 *head) {
    int count = 0;
    const struct Node668_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
