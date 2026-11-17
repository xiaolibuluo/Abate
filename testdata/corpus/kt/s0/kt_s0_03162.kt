package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25970<A, B>(val first: A, val second: B)

fun combine25970(left: Int, right: Int): Pair25970<Int, Int> {
    return Pair25970(left + 10, right + 10)
}
fun <T> List<T>.secondOrNull25971(): T? = if (size > 1) this[1] else null

fun render25971(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25972<A, B>(val first: A, val second: B)

fun combine25972(left: Int, right: Int): Pair25972<Int, Int> {
    return Pair25972(left + 24, right + 32)
}
fun <T, R> List<T>.mapChecked25973(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry25974 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25974 = "registry-38"
    }
}
