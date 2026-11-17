package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7410<A, B>(val first: A, val second: B)

fun combine7410(left: Int, right: Int): Pair7410<Int, Int> {
    return Pair7410(left + 35, right + 11)
}
data class Pair7411<A, B>(val first: A, val second: B)

fun combine7411(left: Int, right: Int): Pair7411<Int, Int> {
    return Pair7411(left + 42, right + 22)
}
fun <T, R> List<T>.mapChecked7412(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry7413 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7413 = "registry-56"
    }
}
