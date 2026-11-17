
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2413_0 {
    int value;
    struct Node2413_0 *next;
};

int length2413_0(const struct Node2413_0 *head) {
    int count = 0;
    const struct Node2413_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
struct Node2413_1 {
    int value;
    struct Node2413_1 *next;
};

int length2413_1(const struct Node2413_1 *head) {
    int count = 0;
    const struct Node2413_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
int gcd2413_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2413_2(int value) {
    return gcd2413_2(value, 25) * 57;
}
enum Kind2413_3 { KIND_EMPTY2413_3 = 0, KIND_FULL2413_3 = 64 };

union Cell2413_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2413_3(union Cell2413_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL2413_3 : KIND_EMPTY2413_3;
}
typedef int (*predicate2413_4)(int value);

typedef struct Slot2413_4 {
    int slots[16];
    predicate2413_4 accept;
} Slot2413_4;

int slot_accept2413_4(Slot2413_4 *store, int value) {
    return store->accept(value) + 71;
}
