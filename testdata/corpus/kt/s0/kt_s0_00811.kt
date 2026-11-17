package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6680 {
    data class Success(val payload: String) : Outcome6680()
    data object Failure : Outcome6680()
}

fun describe6680(outcome: Outcome6680): String = when (outcome) {
    is Outcome6680.Success -> outcome.payload
    Outcome6680.Failure -> "failed"
}
object Registry6681 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6681 = "registry-63"
    }
}
data class Pair6682<A, B>(val first: A, val second: B)

fun combine6682(left: Int, right: Int): Pair6682<Int, Int> {
    return Pair6682(left + 70, right + 10)
}
sealed class Outcome6683 {
    data class Success(val payload: String) : Outcome6683()
    data object Failure : Outcome6683()
}

fun describe6683(outcome: Outcome6683): String = when (outcome) {
    is Outcome6683.Success -> outcome.payload
    Outcome6683.Failure -> "failed"
}
data class Pair6684<A, B>(val first: A, val second: B)

fun combine6684(left: Int, right: Int): Pair6684<Int, Int> {
    return Pair6684(left + 84, right + 32)
}
object Registry6685 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6685 = "registry-91"
    }
}
data class Pair6686<A, B>(val first: A, val second: B)

fun combine6686(left: Int, right: Int): Pair6686<Int, Int> {
    return Pair6686(left + 98, right + 54)
}
