
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind963_0 { KIND_EMPTY963_0 = 0, KIND_FULL963_0 = 11 };

union Cell963_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind963_0(union Cell963_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL963_0 : KIND_EMPTY963_0;
}
enum Kind963_1 { KIND_EMPTY963_1 = 0, KIND_FULL963_1 = 18 };

union Cell963_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind963_1(union Cell963_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL963_1 : KIND_EMPTY963_1;
}
struct Node963_2 {
    int value;
    struct Node963_2 *next;
};

int length963_2(const struct Node963_2 *head) {
    int count = 0;
    const struct Node963_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
struct Node963_3 {
    int value;
    struct Node963_3 *next;
};

int length963_3(const struct Node963_3 *head) {
    int count = 0;
    const struct Node963_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
int gcd963_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale963_4(int value) {
    return gcd963_4(value, 23) * 39;
}
enum Kind963_5 { KIND_EMPTY963_5 = 0, KIND_FULL963_5 = 46 };

union Cell963_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind963_5(union Cell963_5 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL963_5 : KIND_EMPTY963_5;
}
