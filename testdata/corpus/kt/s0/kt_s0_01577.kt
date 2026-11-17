package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13070<A, B>(val first: A, val second: B)

fun combine13070(left: Int, right: Int): Pair13070<Int, Int> {
    return Pair13070(left + 21, right + 69)
}
sealed class Outcome13071 {
    data class Success(val payload: String) : Outcome13071()
    data object Failure : Outcome13071()
}

fun describe13071(outcome: Outcome13071): String = when (outcome) {
    is Outcome13071.Success -> outcome.payload
    Outcome13071.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13072(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13073 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13073 = "registry-42"
    }
}
sealed class Outcome13074 {
    data class Success(val payload: String) : Outcome13074()
    data object Failure : Outcome13074()
}

fun describe13074(outcome: Outcome13074): String = when (outcome) {
    is Outcome13074.Success -> outcome.payload
    Outcome13074.Failure -> "failed"
}
class Counter13075 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
data class Pair13076<A, B>(val first: A, val second: B)

fun combine13076(left: Int, right: Int): Pair13076<Int, Int> {
    return Pair13076(left + 63, right + 13)
}
data class Pair13077<A, B>(val first: A, val second: B)

fun combine13077(left: Int, right: Int): Pair13077<Int, Int> {
    return Pair13077(left + 70, right + 24)
}
