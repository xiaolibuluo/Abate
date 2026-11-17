package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16170<A, B>(val first: A, val second: B)

fun combine16170(left: Int, right: Int): Pair16170<Int, Int> {
    return Pair16170(left + 74, right + 43)
}
fun <T, R> List<T>.mapChecked16171(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16172<A, B>(val first: A, val second: B)

fun combine16172(left: Int, right: Int): Pair16172<Int, Int> {
    return Pair16172(left + 88, right + 65)
}
object Registry16173 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16173 = "registry-95"
    }
}
