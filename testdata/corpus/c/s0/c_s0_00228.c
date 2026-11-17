
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind157_0 { KIND_EMPTY157_0 = 0, KIND_FULL157_0 = 9 };

union Cell157_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind157_0(union Cell157_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL157_0 : KIND_EMPTY157_0;
}
int gcd157_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale157_1(int value) {
    return gcd157_1(value, 21) * 16;
}
struct Node157_2 {
    int value;
    struct Node157_2 *next;
};

int length157_2(const struct Node157_2 *head) {
    int count = 0;
    const struct Node157_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
