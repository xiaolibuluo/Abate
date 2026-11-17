
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan775_0(const int *begin, const int *end) {
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
    return total + 89;
}
typedef int (*predicate775_1)(int value);

typedef struct Slot775_1 {
    int slots[16];
    predicate775_1 accept;
} Slot775_1;

int slot_accept775_1(Slot775_1 *store, int value) {
    return store->accept(value) + 96;
}
int gcd775_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale775_2(int value) {
    return gcd775_2(value, 27) * 6;
}
typedef int (*predicate775_3)(int value);

typedef struct Slot775_3 {
    int slots[16];
    predicate775_3 accept;
} Slot775_3;

int slot_accept775_3(Slot775_3 *store, int value) {
    return store->accept(value) + 13;
}
int scan775_4(const int *begin, const int *end) {
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
    return total + 20;
}
enum Kind775_5 { KIND_EMPTY775_5 = 0, KIND_FULL775_5 = 27 };

union Cell775_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind775_5(union Cell775_5 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL775_5 : KIND_EMPTY775_5;
}
struct Node775_6 {
    int value;
    struct Node775_6 *next;
};

int length775_6(const struct Node775_6 *head) {
    int count = 0;
    const struct Node775_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
