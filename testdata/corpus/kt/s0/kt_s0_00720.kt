package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5940 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5940 = "registry-64"
    }
}
fun <T> List<T>.secondOrNull5941(): T? = if (size > 1) this[1] else null

fun render5941(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5942<A, B>(val first: A, val second: B)

fun combine5942(left: Int, right: Int): Pair5942<Int, Int> {
    return Pair5942(left + 78, right + 41)
}
