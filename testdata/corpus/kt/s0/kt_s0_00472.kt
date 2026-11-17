package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3860(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter3861 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T> List<T>.secondOrNull3862(): T? = if (size > 1) this[1] else null

fun render3862(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3863<A, B>(val first: A, val second: B)

fun combine3863(left: Int, right: Int): Pair3863<Int, Int> {
    return Pair3863(left + 97, right + 60)
}
@JvmName("step3864")
fun step3864(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
data class Pair3865<A, B>(val first: A, val second: B)

fun combine3865(left: Int, right: Int): Pair3865<Int, Int> {
    return Pair3865(left + 14, right + 21)
}
object Registry3866 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3866 = "registry-21"
    }
}
fun <T> List<T>.secondOrNull3867(): T? = if (size > 1) this[1] else null

fun render3867(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3868 {
    data class Success(val payload: String) : Outcome3868()
    data object Failure : Outcome3868()
}

fun describe3868(outcome: Outcome3868): String = when (outcome) {
    is Outcome3868.Success -> outcome.payload
    Outcome3868.Failure -> "failed"
}
