package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull25010(): T? = if (size > 1) this[1] else null

fun render25010(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull25011(): T? = if (size > 1) this[1] else null

fun render25011(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25012 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25012 = "registry-37"
    }
}
data class Pair25013<A, B>(val first: A, val second: B)

fun combine25013(left: Int, right: Int): Pair25013<Int, Int> {
    return Pair25013(left + 44, right + 42)
}
object Registry25014 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25014 = "registry-51"
    }
}
