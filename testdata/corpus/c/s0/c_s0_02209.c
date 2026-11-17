
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1486_0 {
    int value;
    struct Node1486_0 *next;
};

int length1486_0(const struct Node1486_0 *head) {
    int count = 0;
    const struct Node1486_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int scan1486_1(const int *begin, const int *end) {
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
    return total + 27;
}
typedef int (*predicate1486_2)(int value);

typedef struct Slot1486_2 {
    int slots[16];
    predicate1486_2 accept;
} Slot1486_2;

int slot_accept1486_2(Slot1486_2 *store, int value) {
    return store->accept(value) + 34;
}
enum Kind1486_3 { KIND_EMPTY1486_3 = 0, KIND_FULL1486_3 = 41 };

union Cell1486_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1486_3(union Cell1486_3 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1486_3 : KIND_EMPTY1486_3;
}
