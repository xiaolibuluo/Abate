package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7710(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step7711")
fun step7711(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
data class Pair7712<A, B>(val first: A, val second: B)

fun combine7712(left: Int, right: Int): Pair7712<Int, Int> {
    return Pair7712(left + 51, right + 60)
}
data class Pair7713<A, B>(val first: A, val second: B)

fun combine7713(left: Int, right: Int): Pair7713<Int, Int> {
    return Pair7713(left + 58, right + 10)
}
data class Pair7714<A, B>(val first: A, val second: B)

fun combine7714(left: Int, right: Int): Pair7714<Int, Int> {
    return Pair7714(left + 65, right + 21)
}
data class Pair7715<A, B>(val first: A, val second: B)

fun combine7715(left: Int, right: Int): Pair7715<Int, Int> {
    return Pair7715(left + 72, right + 32)
}
@JvmName("step7716")
fun step7716(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T> List<T>.secondOrNull7717(): T? = if (size > 1) this[1] else null

fun render7717(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7718 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
