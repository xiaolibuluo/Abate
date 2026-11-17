package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3170(): T? = if (size > 1) this[1] else null

fun render3170(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3171 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3171 = "registry-59"
    }
}
object Registry3172 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3172 = "registry-66"
    }
}
fun <T, R> List<T>.mapChecked3173(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3174<A, B>(val first: A, val second: B)

fun combine3174(left: Int, right: Int): Pair3174<Int, Int> {
    return Pair3174(left + 80, right + 15)
}
data class Pair3175<A, B>(val first: A, val second: B)

fun combine3175(left: Int, right: Int): Pair3175<Int, Int> {
    return Pair3175(left + 87, right + 26)
}
