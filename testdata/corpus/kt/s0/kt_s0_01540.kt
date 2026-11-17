package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12760(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair12761<A, B>(val first: A, val second: B)

fun combine12761(left: Int, right: Int): Pair12761<Int, Int> {
    return Pair12761(left + 13, right + 46)
}
object Registry12762 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12762 = "registry-20"
    }
}
