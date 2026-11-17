package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19800<A, B>(val first: A, val second: B)

fun combine19800(left: Int, right: Int): Pair19800<Int, Int> {
    return Pair19800(left + 40, right + 22)
}
sealed class Outcome19801 {
    data class Success(val payload: String) : Outcome19801()
    data object Failure : Outcome19801()
}

fun describe19801(outcome: Outcome19801): String = when (outcome) {
    is Outcome19801.Success -> outcome.payload
    Outcome19801.Failure -> "failed"
}
object Registry19802 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19802 = "registry-54"
    }
}
fun <T, R> List<T>.mapChecked19803(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19804<A, B>(val first: A, val second: B)

fun combine19804(left: Int, right: Int): Pair19804<Int, Int> {
    return Pair19804(left + 68, right + 66)
}
sealed class Outcome19805 {
    data class Success(val payload: String) : Outcome19805()
    data object Failure : Outcome19805()
}

fun describe19805(outcome: Outcome19805): String = when (outcome) {
    is Outcome19805.Success -> outcome.payload
    Outcome19805.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19806(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
