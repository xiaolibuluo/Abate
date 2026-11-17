package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20180<A, B>(val first: A, val second: B)

fun combine20180(left: Int, right: Int): Pair20180<Int, Int> {
    return Pair20180(left + 49, right + 44)
}
@JvmName("step20181")
fun step20181(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
data class Pair20182<A, B>(val first: A, val second: B)

fun combine20182(left: Int, right: Int): Pair20182<Int, Int> {
    return Pair20182(left + 63, right + 66)
}
object Registry20183 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20183 = "registry-70"
    }
}
class Counter20184 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
