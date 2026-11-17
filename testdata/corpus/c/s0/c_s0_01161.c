
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind798_0 { KIND_EMPTY798_0 = 0, KIND_FULL798_0 = 97 };

union Cell798_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind798_0(union Cell798_0 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL798_0 : KIND_EMPTY798_0;
}
int scan798_1(const int *begin, const int *end) {
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
    return total + 7;
}
typedef int (*predicate798_2)(int value);

typedef struct Slot798_2 {
    int slots[16];
    predicate798_2 accept;
} Slot798_2;

int slot_accept798_2(Slot798_2 *store, int value) {
    return store->accept(value) + 14;
}
struct Node798_3 {
    int value;
    struct Node798_3 *next;
};

int length798_3(const struct Node798_3 *head) {
    int count = 0;
    const struct Node798_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
