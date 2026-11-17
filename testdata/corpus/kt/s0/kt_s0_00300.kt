package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2510(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2511<A, B>(val first: A, val second: B)

fun combine2511(left: Int, right: Int): Pair2511<Int, Int> {
    return Pair2511(left + 74, right + 69)
}
object Registry2512 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2512 = "registry-81"
    }
}
