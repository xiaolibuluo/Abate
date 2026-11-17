package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24950(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24951 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24951 = "registry-49"
    }
}
data class Pair24952<A, B>(val first: A, val second: B)

fun combine24952(left: Int, right: Int): Pair24952<Int, Int> {
    return Pair24952(left + 56, right + 50)
}
fun <T> List<T>.secondOrNull24953(): T? = if (size > 1) this[1] else null

fun render24953(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
