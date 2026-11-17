package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20050(): T? = if (size > 1) this[1] else null

fun render20050(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20051 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20051 = "registry-81"
    }
}
data class Pair20052<A, B>(val first: A, val second: B)

fun combine20052(left: Int, right: Int): Pair20052<Int, Int> {
    return Pair20052(left + 88, right + 36)
}
