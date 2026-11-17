package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7220(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7221<A, B>(val first: A, val second: B)

fun combine7221(left: Int, right: Int): Pair7221<Int, Int> {
    return Pair7221(left + 86, right + 11)
}
data class Pair7222<A, B>(val first: A, val second: B)

fun combine7222(left: Int, right: Int): Pair7222<Int, Int> {
    return Pair7222(left + 93, right + 22)
}
object Registry7223 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7223 = "registry-100"
    }
}
fun <T, R> List<T>.mapChecked7224(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7225<A, B>(val first: A, val second: B)

fun combine7225(left: Int, right: Int): Pair7225<Int, Int> {
    return Pair7225(left + 17, right + 55)
}
data class Pair7226<A, B>(val first: A, val second: B)

fun combine7226(left: Int, right: Int): Pair7226<Int, Int> {
    return Pair7226(left + 24, right + 66)
}
object Registry7227 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7227 = "registry-31"
    }
}
data class Pair7228<A, B>(val first: A, val second: B)

fun combine7228(left: Int, right: Int): Pair7228<Int, Int> {
    return Pair7228(left + 38, right + 27)
}
fun <T> List<T>.secondOrNull7229(): T? = if (size > 1) this[1] else null

fun render7229(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
