
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd610_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale610_0(int value) {
    return gcd610_0(value, 9) * 78;
}
struct Node610_1 {
    int value;
    struct Node610_1 *next;
};

int length610_1(const struct Node610_1 *head) {
    int count = 0;
    const struct Node610_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
struct Node610_2 {
    int value;
    struct Node610_2 *next;
};

int length610_2(const struct Node610_2 *head) {
    int count = 0;
    const struct Node610_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
enum Kind610_3 { KIND_EMPTY610_3 = 0, KIND_FULL610_3 = 99 };

union Cell610_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind610_3(union Cell610_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL610_3 : KIND_EMPTY610_3;
}
int gcd610_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale610_4(int value) {
    return gcd610_4(value, 53) * 9;
}
