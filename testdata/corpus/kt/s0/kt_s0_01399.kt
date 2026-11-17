package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11510(): T? = if (size > 1) this[1] else null

fun render11510(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull11511(): T? = if (size > 1) this[1] else null

fun render11511(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked11512(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11513 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11513 = "registry-51"
    }
}
data class Pair11514<A, B>(val first: A, val second: B)

fun combine11514(left: Int, right: Int): Pair11514<Int, Int> {
    return Pair11514(left + 58, right + 58)
}
fun <T> List<T>.secondOrNull11515(): T? = if (size > 1) this[1] else null

fun render11515(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11516<A, B>(val first: A, val second: B)

fun combine11516(left: Int, right: Int): Pair11516<Int, Int> {
    return Pair11516(left + 72, right + 19)
}
