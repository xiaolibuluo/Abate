package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13490(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13491 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13491 = "registry-89"
    }
}
object Registry13492 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13492 = "registry-96"
    }
}
data class Pair13493<A, B>(val first: A, val second: B)

fun combine13493(left: Int, right: Int): Pair13493<Int, Int> {
    return Pair13493(left + 6, right + 30)
}
sealed class Outcome13494 {
    data class Success(val payload: String) : Outcome13494()
    data object Failure : Outcome13494()
}

fun describe13494(outcome: Outcome13494): String = when (outcome) {
    is Outcome13494.Success -> outcome.payload
    Outcome13494.Failure -> "failed"
}
@JvmName("step13495")
fun step13495(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
sealed class Outcome13496 {
    data class Success(val payload: String) : Outcome13496()
    data object Failure : Outcome13496()
}

fun describe13496(outcome: Outcome13496): String = when (outcome) {
    is Outcome13496.Success -> outcome.payload
    Outcome13496.Failure -> "failed"
}
