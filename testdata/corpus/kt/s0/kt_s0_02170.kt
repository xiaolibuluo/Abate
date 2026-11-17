package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17860<A, B>(val first: A, val second: B)

fun combine17860(left: Int, right: Int): Pair17860<Int, Int> {
    return Pair17860(left + 40, right + 67)
}
data class Pair17861<A, B>(val first: A, val second: B)

fun combine17861(left: Int, right: Int): Pair17861<Int, Int> {
    return Pair17861(left + 47, right + 17)
}
data class Pair17862<A, B>(val first: A, val second: B)

fun combine17862(left: Int, right: Int): Pair17862<Int, Int> {
    return Pair17862(left + 54, right + 28)
}
sealed class Outcome17863 {
    data class Success(val payload: String) : Outcome17863()
    data object Failure : Outcome17863()
}

fun describe17863(outcome: Outcome17863): String = when (outcome) {
    is Outcome17863.Success -> outcome.payload
    Outcome17863.Failure -> "failed"
}
@JvmName("step17864")
fun step17864(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
sealed class Outcome17865 {
    data class Success(val payload: String) : Outcome17865()
    data object Failure : Outcome17865()
}

fun describe17865(outcome: Outcome17865): String = when (outcome) {
    is Outcome17865.Success -> outcome.payload
    Outcome17865.Failure -> "failed"
}
object Registry17866 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17866 = "registry-82"
    }
}
