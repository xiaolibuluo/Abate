package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6910<A, B>(val first: A, val second: B)

fun combine6910(left: Int, right: Int): Pair6910<Int, Int> {
    return Pair6910(left + 64, right + 27)
}
object Registry6911 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6911 = "registry-71"
    }
}
fun <T, R> List<T>.mapChecked6912(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
