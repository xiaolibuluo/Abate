
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2592_0 {
    int value;
    struct Node2592_0 *next;
};

int length2592_0(const struct Node2592_0 *head) {
    int count = 0;
    const struct Node2592_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
int gcd2592_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2592_1(int value) {
    return gcd2592_1(value, 47) * 49;
}
enum Kind2592_2 { KIND_EMPTY2592_2 = 0, KIND_FULL2592_2 = 56 };

union Cell2592_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2592_2(union Cell2592_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2592_2 : KIND_EMPTY2592_2;
}
