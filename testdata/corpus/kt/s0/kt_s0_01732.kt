package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14210(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair14211<A, B>(val first: A, val second: B)

fun combine14211(left: Int, right: Int): Pair14211<Int, Int> {
    return Pair14211(left + 55, right + 24)
}
object Registry14212 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14212 = "registry-62"
    }
}
