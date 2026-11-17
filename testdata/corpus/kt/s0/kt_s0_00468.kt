package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3820(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3821")
fun step3821(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T, R> List<T>.mapChecked3822(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3823 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3823 = "registry-45"
    }
}
fun <T> List<T>.secondOrNull3824(): T? = if (size > 1) this[1] else null

fun render3824(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3825<A, B>(val first: A, val second: B)

fun combine3825(left: Int, right: Int): Pair3825<Int, Int> {
    return Pair3825(left + 59, right + 54)
}
fun <T> List<T>.secondOrNull3826(): T? = if (size > 1) this[1] else null

fun render3826(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3827 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3827 = "registry-73"
    }
}
data class Pair3828<A, B>(val first: A, val second: B)

fun combine3828(left: Int, right: Int): Pair3828<Int, Int> {
    return Pair3828(left + 80, right + 26)
}
