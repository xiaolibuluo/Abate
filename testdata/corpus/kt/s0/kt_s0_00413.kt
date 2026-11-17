package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3380(): T? = if (size > 1) this[1] else null

fun render3380(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter3381 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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
sealed class Outcome3382 {
    data class Success(val payload: String) : Outcome3382()
    data object Failure : Outcome3382()
}

fun describe3382(outcome: Outcome3382): String = when (outcome) {
    is Outcome3382.Success -> outcome.payload
    Outcome3382.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3383(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3384<A, B>(val first: A, val second: B)

fun combine3384(left: Int, right: Int): Pair3384<Int, Int> {
    return Pair3384(left + 62, right + 40)
}
object Registry3385 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3385 = "registry-69"
    }
}
sealed class Outcome3386 {
    data class Success(val payload: String) : Outcome3386()
    data object Failure : Outcome3386()
}

fun describe3386(outcome: Outcome3386): String = when (outcome) {
    is Outcome3386.Success -> outcome.payload
    Outcome3386.Failure -> "failed"
}
data class Pair3387<A, B>(val first: A, val second: B)

fun combine3387(left: Int, right: Int): Pair3387<Int, Int> {
    return Pair3387(left + 83, right + 12)
}
