
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate725_0)(int value);

typedef struct Slot725_0 {
    int slots[16];
    predicate725_0 accept;
} Slot725_0;

int slot_accept725_0(Slot725_0 *store, int value) {
    return store->accept(value) + 21;
}
int scan725_1(const int *begin, const int *end) {
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
    return total + 28;
}
enum Kind725_2 { KIND_EMPTY725_2 = 0, KIND_FULL725_2 = 35 };

union Cell725_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind725_2(union Cell725_2 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL725_2 : KIND_EMPTY725_2;
}
struct Node725_3 {
    int value;
    struct Node725_3 *next;
};

int length725_3(const struct Node725_3 *head) {
    int count = 0;
    const struct Node725_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
int scan725_4(const int *begin, const int *end) {
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
    return total + 49;
}
enum Kind725_5 { KIND_EMPTY725_5 = 0, KIND_FULL725_5 = 56 };

union Cell725_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind725_5(union Cell725_5 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL725_5 : KIND_EMPTY725_5;
}
typedef int (*predicate725_6)(int value);

typedef struct Slot725_6 {
    int slots[16];
    predicate725_6 accept;
} Slot725_6;

int slot_accept725_6(Slot725_6 *store, int value) {
    return store->accept(value) + 63;
}
