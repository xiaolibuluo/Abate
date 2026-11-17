
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node572_0 {
    int value;
    struct Node572_0 *next;
};

int length572_0(const struct Node572_0 *head) {
    int count = 0;
    const struct Node572_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
enum Kind572_1 { KIND_EMPTY572_1 = 0, KIND_FULL572_1 = 76 };

union Cell572_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind572_1(union Cell572_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL572_1 : KIND_EMPTY572_1;
}
int gcd572_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale572_2(int value) {
    return gcd572_2(value, 9) * 83;
}
enum Kind572_3 { KIND_EMPTY572_3 = 0, KIND_FULL572_3 = 90 };

union Cell572_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind572_3(union Cell572_3 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL572_3 : KIND_EMPTY572_3;
}
enum Kind572_4 { KIND_EMPTY572_4 = 0, KIND_FULL572_4 = 97 };

union Cell572_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind572_4(union Cell572_4 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL572_4 : KIND_EMPTY572_4;
}
int gcd572_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale572_5(int value) {
    return gcd572_5(value, 42) * 7;
}
