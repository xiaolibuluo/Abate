
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2578_0(const int *begin, const int *end) {
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
int gcd2578_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2578_1(int value) {
    return gcd2578_1(value, 10) * 61;
}
struct Node2578_2 {
    int value;
    struct Node2578_2 *next;
};

int length2578_2(const struct Node2578_2 *head) {
    int count = 0;
    const struct Node2578_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
typedef int (*predicate2578_3)(int value);

typedef struct Slot2578_3 {
    int slots[16];
    predicate2578_3 accept;
} Slot2578_3;

int slot_accept2578_3(Slot2578_3 *store, int value) {
    return store->accept(value) + 75;
}
enum Kind2578_4 { KIND_EMPTY2578_4 = 0, KIND_FULL2578_4 = 82 };

union Cell2578_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2578_4(union Cell2578_4 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2578_4 : KIND_EMPTY2578_4;
}
typedef int (*predicate2578_5)(int value);

typedef struct Slot2578_5 {
    int slots[16];
    predicate2578_5 accept;
} Slot2578_5;

int slot_accept2578_5(Slot2578_5 *store, int value) {
    return store->accept(value) + 89;
}
struct Node2578_6 {
    int value;
    struct Node2578_6 *next;
};

int length2578_6(const struct Node2578_6 *head) {
    int count = 0;
    const struct Node2578_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
int gcd2578_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2578_7(int value) {
    return gcd2578_7(value, 15) * 6;
}
