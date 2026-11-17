package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG580 = "registry-80"
    }
}
fun <T> List<T>.secondOrNull581(): T? = if (size > 1) this[1] else null

fun render581(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair582<A, B>(val first: A, val second: B)

fun combine582(left: Int, right: Int): Pair582<Int, Int> {
    return Pair582(left + 94, right + 10)
}
