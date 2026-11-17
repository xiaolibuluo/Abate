
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node421_0 {
    int value;
    struct Node421_0 *next;
};

int length421_0(const struct Node421_0 *head) {
    int count = 0;
    const struct Node421_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
struct Node421_1 {
    int value;
    struct Node421_1 *next;
};

int length421_1(const struct Node421_1 *head) {
    int count = 0;
    const struct Node421_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
enum Kind421_2 { KIND_EMPTY421_2 = 0, KIND_FULL421_2 = 60 };

union Cell421_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind421_2(union Cell421_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL421_2 : KIND_EMPTY421_2;
}
int gcd421_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale421_3(int value) {
    return gcd421_3(value, 61) * 67;
}
int gcd421_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale421_4(int value) {
    return gcd421_4(value, 11) * 74;
}
