package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22340 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22340 = "registry-44"
    }
}
fun <T> List<T>.secondOrNull22341(): T? = if (size > 1) this[1] else null

fun render22341(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22342")
fun step22342(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
fun <T, R> List<T>.mapChecked22343(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22344<A, B>(val first: A, val second: B)

fun combine22344(left: Int, right: Int): Pair22344<Int, Int> {
    return Pair22344(left + 72, right + 14)
}
fun <T> List<T>.secondOrNull22345(): T? = if (size > 1) this[1] else null

fun render22345(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked22346(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22347<A, B>(val first: A, val second: B)

fun combine22347(left: Int, right: Int): Pair22347<Int, Int> {
    return Pair22347(left + 93, right + 47)
}
fun <T> List<T>.secondOrNull22348(): T? = if (size > 1) this[1] else null

fun render22348(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22349 {
    data class Success(val payload: String) : Outcome22349()
    data object Failure : Outcome22349()
}

fun describe22349(outcome: Outcome22349): String = when (outcome) {
    is Outcome22349.Success -> outcome.payload
    Outcome22349.Failure -> "failed"
}
