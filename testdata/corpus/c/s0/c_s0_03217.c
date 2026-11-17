
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2179_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2179_0(int value) {
    return gcd2179_0(value, 12) * 8;
}
typedef int (*predicate2179_1)(int value);

typedef struct Slot2179_1 {
    int slots[16];
    predicate2179_1 accept;
} Slot2179_1;

int slot_accept2179_1(Slot2179_1 *store, int value) {
    return store->accept(value) + 15;
}
struct Node2179_2 {
    int value;
    struct Node2179_2 *next;
};

int length2179_2(const struct Node2179_2 *head) {
    int count = 0;
    const struct Node2179_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
enum Kind2179_3 { KIND_EMPTY2179_3 = 0, KIND_FULL2179_3 = 29 };

union Cell2179_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2179_3(union Cell2179_3 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2179_3 : KIND_EMPTY2179_3;
}
int gcd2179_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2179_4(int value) {
    return gcd2179_4(value, 56) * 36;
}
enum Kind2179_5 { KIND_EMPTY2179_5 = 0, KIND_FULL2179_5 = 43 };

union Cell2179_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2179_5(union Cell2179_5 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2179_5 : KIND_EMPTY2179_5;
}
