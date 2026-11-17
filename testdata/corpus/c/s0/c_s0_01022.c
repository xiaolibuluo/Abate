
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan701_0(const int *begin, const int *end) {
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
    return total + 97;
}
typedef int (*predicate701_1)(int value);

typedef struct Slot701_1 {
    int slots[16];
    predicate701_1 accept;
} Slot701_1;

int slot_accept701_1(Slot701_1 *store, int value) {
    return store->accept(value) + 7;
}
enum Kind701_2 { KIND_EMPTY701_2 = 0, KIND_FULL701_2 = 14 };

union Cell701_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind701_2(union Cell701_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL701_2 : KIND_EMPTY701_2;
}
struct Node701_3 {
    int value;
    struct Node701_3 *next;
};

int length701_3(const struct Node701_3 *head) {
    int count = 0;
    const struct Node701_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
