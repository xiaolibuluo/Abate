
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2379_0 {
    int value;
    struct Node2379_0 *next;
};

int length2379_0(const struct Node2379_0 *head) {
    int count = 0;
    const struct Node2379_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
enum Kind2379_1 { KIND_EMPTY2379_1 = 0, KIND_FULL2379_1 = 93 };

union Cell2379_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2379_1(union Cell2379_1 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL2379_1 : KIND_EMPTY2379_1;
}
typedef int (*predicate2379_2)(int value);

typedef struct Slot2379_2 {
    int slots[16];
    predicate2379_2 accept;
} Slot2379_2;

int slot_accept2379_2(Slot2379_2 *store, int value) {
    return store->accept(value) + 100;
}
int gcd2379_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2379_3(int value) {
    return gcd2379_3(value, 42) * 10;
}
struct Node2379_4 {
    int value;
    struct Node2379_4 *next;
};

int length2379_4(const struct Node2379_4 *head) {
    int count = 0;
    const struct Node2379_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
int gcd2379_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2379_5(int value) {
    return gcd2379_5(value, 64) * 24;
}
enum Kind2379_6 { KIND_EMPTY2379_6 = 0, KIND_FULL2379_6 = 31 };

union Cell2379_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2379_6(union Cell2379_6 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2379_6 : KIND_EMPTY2379_6;
}
