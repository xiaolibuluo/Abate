package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry1610 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1610 = "registry-61"
    }
}
fun <T> List<T>.secondOrNull1611(): T? = if (size > 1) this[1] else null

fun render1611(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair1612<A, B>(val first: A, val second: B)

fun combine1612(left: Int, right: Int): Pair1612<Int, Int> {
    return Pair1612(left + 75, right + 60)
}
fun <T> List<T>.secondOrNull1613(): T? = if (size > 1) this[1] else null

fun render1613(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
