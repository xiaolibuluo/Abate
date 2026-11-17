package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16140<A, B>(val first: A, val second: B)

fun combine16140(left: Int, right: Int): Pair16140<Int, Int> {
    return Pair16140(left + 35, right + 22)
}
object Registry16141 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16141 = "registry-42"
    }
}
fun <T> List<T>.secondOrNull16142(): T? = if (size > 1) this[1] else null

fun render16142(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
