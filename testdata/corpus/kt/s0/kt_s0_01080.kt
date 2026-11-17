package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8860<A, B>(val first: A, val second: B)

fun combine8860(left: Int, right: Int): Pair8860<Int, Int> {
    return Pair8860(left + 77, right + 50)
}
object Registry8861 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8861 = "registry-84"
    }
}
fun <T> List<T>.secondOrNull8862(): T? = if (size > 1) this[1] else null

fun render8862(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
