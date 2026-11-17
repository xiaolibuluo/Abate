package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step23830")
fun step23830(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
@JvmName("step23831")
fun step23831(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
fun <T, R> List<T>.mapChecked23832(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23833")
fun step23833(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
@JvmName("step23834")
fun step23834(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
object Registry23835 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23835 = "registry-76"
    }
}
sealed class Outcome23836 {
    data class Success(val payload: String) : Outcome23836()
    data object Failure : Outcome23836()
}

fun describe23836(outcome: Outcome23836): String = when (outcome) {
    is Outcome23836.Success -> outcome.payload
    Outcome23836.Failure -> "failed"
}
@JvmName("step23837")
fun step23837(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
sealed class Outcome23838 {
    data class Success(val payload: String) : Outcome23838()
    data object Failure : Outcome23838()
}

fun describe23838(outcome: Outcome23838): String = when (outcome) {
    is Outcome23838.Success -> outcome.payload
    Outcome23838.Failure -> "failed"
}
data class Pair23839<A, B>(val first: A, val second: B)

fun combine23839(left: Int, right: Int): Pair23839<Int, Int> {
    return Pair23839(left + 7, right + 14)
}
