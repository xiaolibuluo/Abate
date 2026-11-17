package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12540<A, B>(val first: A, val second: B)

fun combine12540(left: Int, right: Int): Pair12540<Int, Int> {
    return Pair12540(left + 11, right + 64)
}
@JvmName("step12541")
fun step12541(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
data class Pair12542<A, B>(val first: A, val second: B)

fun combine12542(left: Int, right: Int): Pair12542<Int, Int> {
    return Pair12542(left + 25, right + 25)
}
sealed class Outcome12543 {
    data class Success(val payload: String) : Outcome12543()
    data object Failure : Outcome12543()
}

fun describe12543(outcome: Outcome12543): String = when (outcome) {
    is Outcome12543.Success -> outcome.payload
    Outcome12543.Failure -> "failed"
}
data class Pair12544<A, B>(val first: A, val second: B)

fun combine12544(left: Int, right: Int): Pair12544<Int, Int> {
    return Pair12544(left + 39, right + 47)
}
data class Pair12545<A, B>(val first: A, val second: B)

fun combine12545(left: Int, right: Int): Pair12545<Int, Int> {
    return Pair12545(left + 46, right + 58)
}
object Registry12546 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12546 = "registry-53"
    }
}
@JvmName("step12547")
fun step12547(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
data class Pair12548<A, B>(val first: A, val second: B)

fun combine12548(left: Int, right: Int): Pair12548<Int, Int> {
    return Pair12548(left + 67, right + 30)
}
