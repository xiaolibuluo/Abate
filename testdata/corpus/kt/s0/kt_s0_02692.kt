package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22240")
fun step22240(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
class Counter22241 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T> List<T>.secondOrNull22242(): T? = if (size > 1) this[1] else null

fun render22242(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22243")
fun step22243(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
@JvmName("step22244")
fun step22244(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
fun <T, R> List<T>.mapChecked22245(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22246 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
