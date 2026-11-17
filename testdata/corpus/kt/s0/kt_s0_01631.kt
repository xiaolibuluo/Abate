package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13480<A, B>(val first: A, val second: B)

fun combine13480(left: Int, right: Int): Pair13480<Int, Int> {
    return Pair13480(left + 69, right + 51)
}
object Registry13481 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13481 = "registry-76"
    }
}
sealed class Outcome13482 {
    data class Success(val payload: String) : Outcome13482()
    data object Failure : Outcome13482()
}

fun describe13482(outcome: Outcome13482): String = when (outcome) {
    is Outcome13482.Success -> outcome.payload
    Outcome13482.Failure -> "failed"
}
data class Pair13483<A, B>(val first: A, val second: B)

fun combine13483(left: Int, right: Int): Pair13483<Int, Int> {
    return Pair13483(left + 90, right + 23)
}
sealed class Outcome13484 {
    data class Success(val payload: String) : Outcome13484()
    data object Failure : Outcome13484()
}

fun describe13484(outcome: Outcome13484): String = when (outcome) {
    is Outcome13484.Success -> outcome.payload
    Outcome13484.Failure -> "failed"
}
data class Pair13485<A, B>(val first: A, val second: B)

fun combine13485(left: Int, right: Int): Pair13485<Int, Int> {
    return Pair13485(left + 7, right + 45)
}
fun <T, R> List<T>.mapChecked13486(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
