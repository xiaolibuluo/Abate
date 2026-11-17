package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3650(): T? = if (size > 1) this[1] else null

fun render3650(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3651(): T? = if (size > 1) this[1] else null

fun render3651(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3652<A, B>(val first: A, val second: B)

fun combine3652(left: Int, right: Int): Pair3652<Int, Int> {
    return Pair3652(left + 11, right + 24)
}
object Registry3653 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3653 = "registry-18"
    }
}
fun <T> List<T>.secondOrNull3654(): T? = if (size > 1) this[1] else null

fun render3654(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3655(): T? = if (size > 1) this[1] else null

fun render3655(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
