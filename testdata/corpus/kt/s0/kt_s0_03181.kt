package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26140 {
    data class Success(val payload: String) : Outcome26140()
    data object Failure : Outcome26140()
}

fun describe26140(outcome: Outcome26140): String = when (outcome) {
    is Outcome26140.Success -> outcome.payload
    Outcome26140.Failure -> "failed"
}
@JvmName("step26141")
fun step26141(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
object Registry26142 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26142 = "registry-51"
    }
}
data class Pair26143<A, B>(val first: A, val second: B)

fun combine26143(left: Int, right: Int): Pair26143<Int, Int> {
    return Pair26143(left + 58, right + 40)
}
@JvmName("step26144")
fun step26144(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
@JvmName("step26145")
fun step26145(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
@JvmName("step26146")
fun step26146(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
object Registry26147 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26147 = "registry-86"
    }
}
