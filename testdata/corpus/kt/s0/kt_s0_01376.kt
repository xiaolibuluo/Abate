package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11340(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11341<A, B>(val first: A, val second: B)

fun combine11341(left: Int, right: Int): Pair11341<Int, Int> {
    return Pair11341(left + 10, right + 28)
}
fun <T> List<T>.secondOrNull11342(): T? = if (size > 1) this[1] else null

fun render11342(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11343 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11343 = "registry-24"
    }
}
